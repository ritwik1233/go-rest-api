import React from "react";
import {
  AppBar,
  Toolbar,
  Typography,
  Button,
  Modal,
  Paper,
  Tabs,
  Tab,
} from "@material-ui/core";
import { connect } from "react-redux";
import axios from "axios";
import PropTypes from "prop-types";
import Grid from "@material-ui/core/Grid";
import { useDispatch } from "react-redux";
import { logoutUser, getCurrentUser } from "../actions/index";
import LoginComponent from "./LoginComponent.js";
import RegisterComponent from "./RegisterComponent.js";

function Header(props) {
  const dispatch = useDispatch();
  const [loginModal, setLoginModal] = React.useState(false);
  const [tabValue, setTabValue] = React.useState(0);
  const handleClose = () => {
    setLoginModal(!loginModal);
  };

  const handleLogin = (userDetails) => {
    console.log(userDetails);
    axios
      .post("/api/login", userDetails, {
        headers: { "Content-Type": "application/json" },
      })
      .then(() => {
        dispatch(getCurrentUser());
        setLoginModal(false);
      })
      .catch((err) => {
        console.log(err);
        alert("Login Failed");
      });
  };
  const handleRegister = (userDetails) => {
    axios
      .post("/api/register", userDetails, {
        headers: { "Content-Type": "application/json" },
      })
      .then(() => {
        setLoginModal(false);
      })
      .catch((err) => {
        console.log(err);
        alert("Registration Failed");
      });
  };
  const handleTabChange = (event, newValue) => {
    setTabValue(newValue);
  };
  const loginComponent = (
    <Modal open={loginModal} onClose={handleClose}>
      <Paper
        style={{
          marginTop: "250px",
          marginLeft: "25%",
          width: "50%",
        }}
      >
        <AppBar position="static">
          <Tabs value={tabValue} onChange={handleTabChange}>
            <Tab label="Login" value={0} />
            <Tab label="Register" value={1} />
          </Tabs>
        </AppBar>
        <Paper role="tabpanel" hidden={tabValue !== 0}>
          <LoginComponent handleLogin={handleLogin} />
        </Paper>
        <Paper role="tabpanel" hidden={tabValue !== 1}>
          <RegisterComponent handleRegister={handleRegister} />
        </Paper>
      </Paper>
    </Modal>
  );
  return (
    <AppBar position="static" style={{ width: "100%" }}>
      <Toolbar>
        <Grid container spacing={3}>
          <Grid item xs={12}></Grid>
          <Grid item xs={10}>
            <Typography variant="h6">Go Demo Forum</Typography>
          </Grid>

          {!props.currentUserToken ? (
            <Grid item xs={2}>
              <Button
                fullWidth
                color="inherit"
                onClick={() => {
                  setLoginModal(true);
                }}
              >
                Login
              </Button>
            </Grid>
          ) : (
            <React.Fragment>
              <Grid item xs={1}>
                <Typography variant="caption" style={{ marginTop: "20px" }}>
                  {props.currentUserToken}
                </Typography>
              </Grid>
              <Grid item xs={1}>
                <Button
                  fullWidth
                  size="small"
                  color="inherit"
                  onClick={() => {
                    dispatch(logoutUser(props.currentUserToken));
                  }}
                >
                  Logout
                </Button>
              </Grid>
            </React.Fragment>
          )}
          <Grid item xs={12}>
            {loginComponent}
          </Grid>
        </Grid>
      </Toolbar>
    </AppBar>
  );
}

function mapStateToProps(state) {
  return {
    currentUserToken: state.auth.currentUserToken,
  };
}
// type checking for props
Header.propTypes = {
  currentUserToken: PropTypes.string,
  openLoginModal: PropTypes.func,
};

// setting default props
Header.defaultProps = {
  currentUserToken: null,
  openLoginModal: () => {},
};
export default connect(mapStateToProps)(Header);
