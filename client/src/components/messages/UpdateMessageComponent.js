import React from "react";
import { Grid, TextField, Button } from "@material-ui/core";
import PropTypes from "prop-types";

function UpdateMessageComponet(props) {
  const handleUpdateMessageSubmit = (e) => {
    e.preventDefault();
    let messageDetails = new FormData();
    messageDetails.append("message", e.target.message.value);
    props.updateMessage(messageDetails);
  };

  return (
    <form
      id="update-message"
      onSubmit={handleUpdateMessageSubmit}
      style={{ marginTop: "10px" }}
    >
      <Grid container spacing={3}>
        <Grid item xs={12}>
          <TextField
            name="message"
            fullWidth
            label="Message"
            variant="outlined"
            required
            defaultValue={props.Message}
          />
        </Grid>
        <Grid item xs={12}>
          <Button color="primary" type="submit" fullWidth variant="contained">
            Update Message
          </Button>
        </Grid>
      </Grid>
    </form>
  );
}

// type checking for props
UpdateMessageComponet.propTypes = {
  Message: PropTypes.string,
  updateMessage: PropTypes.func,
};

// setting default props
UpdateMessageComponet.defaultProps = {
  Message: "",
  updateMessage: () => {},
};

export default UpdateMessageComponet;
