import React from "react";
import { Grid, TextField, Button } from "@material-ui/core";
import PropTypes from "prop-types";

function CreateMessageComponent(props) {
  const handleCreateMessageSubmit = (e) => {
    e.preventDefault();
    let messageDetails = {};
    messageDetails.message = e.target.message.value;
    props.handleMessage(messageDetails);
  };

  return (
    <form
      id="create-message"
      onSubmit={handleCreateMessageSubmit}
      style={{ marginTop: "10px" }}
    >
      <Grid container spacing={3}>
        <Grid item xs={1}></Grid>
        <Grid item xs={10}>
          <TextField
            name="message"
            fullWidth
            label="Message"
            variant="outlined"
            required
          />
        </Grid>
        <Grid item xs={1}></Grid>
        <Grid item xs={1}></Grid>
        <Grid item xs={10}>
          <Button color="primary" type="submit" fullWidth variant="contained">
            Create Message
          </Button>
        </Grid>
        <Grid item xs={1}></Grid>
      </Grid>
    </form>
  );
}

// type checking for props
CreateMessageComponent.propTypes = {
  handleMessage: PropTypes.func,
};

// setting default props
CreateMessageComponent.defaultProps = {
  handleMessage: () => {},
};

export default CreateMessageComponent;
