import React from "react";
import { Grid, TextField, Button } from "@material-ui/core";
import PropTypes from "prop-types";

function RegisterComponent(props) {
  const handleRegisterSubmit = (e) => {
    e.preventDefault();
    const registerDetails = new FormData();
    registerDetails.append("email", e.target.email.value);
    registerDetails.append("password", e.target.password.value);
    props.handleRegister(registerDetails);
  };
  return (
    <form
      onSubmit={handleRegisterSubmit}
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
            type="password"
            label="Password"
            variant="outlined"
            required
          />
        </Grid>
        <Grid item xs={1}></Grid>
        <Grid item xs={1}></Grid>
        <Grid item xs={10}>
          <Button color="primary" type="submit" fullWidth variant="contained">
            Register
          </Button>
        </Grid>
        <Grid item xs={1}></Grid>
      </Grid>
    </form>
  );
}

// type checking for props
RegisterComponent.propTypes = {
  handleRegister: PropTypes.func,
};

// setting default props
RegisterComponent.defaultProps = {
  handleRegister: () => {},
};

export default RegisterComponent;
