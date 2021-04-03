import React from "react";
import { Grid, TextField, Button } from "@material-ui/core";
import PropTypes from "prop-types";

function LoginComponent(props) {
  const handleLoginSubmit = (e) => {
    e.preventDefault();
    let loginDetails = new FormData();
    loginDetails.append("email", e.target.email.value);
    loginDetails.append("password", e.target.password.value);
    props.handleLogin(loginDetails);
  };
  return (
    <form
      onSubmit={handleLoginSubmit}
      style={{ marginTop: "10px", width: "inherit" }}
    >
      <Grid container spacing={3}>
        <Grid item xs={1}></Grid>
        <Grid item xs={10}>
          <TextField
            name="email"
            fullWidth
            label="Email"
            variant="outlined"
            required
          />
        </Grid>
        <Grid item xs={1}></Grid>
        <Grid item xs={1}></Grid>
        <Grid item xs={10}>
          <TextField
            name="password"
            fullWidth
            label="Password"
            type="password"
            variant="outlined"
            required
          />
        </Grid>
        <Grid item xs={1}></Grid>
        <Grid item xs={1}></Grid>
        <Grid item xs={10}>
          <Button color="primary" type="submit" fullWidth variant="contained">
            Login
          </Button>
        </Grid>
        <Grid item xs={1}></Grid>
      </Grid>
    </form>
  );
}

// type checking for props
LoginComponent.propTypes = {
  handleLogin: PropTypes.func,
};

// setting default props
LoginComponent.defaultProps = {
  handleLogin: () => {},
};

export default LoginComponent;
