import React from "react";
import { Grid, TextField, Button } from "@material-ui/core";
import PropTypes from "prop-types";

function UpdateCommentComponet(props) {
  const handleCommentSubmit = (e) => {
    e.preventDefault();
    let commentDetails = new FormData();
    commentDetails.append("comment", e.target.message.value);
    props.updateComment(commentDetails, props.commentId);
  };
  return (
    <form
      id="update-comment"
      onSubmit={handleCommentSubmit}
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
            defaultValue={props.Comment}
          />
        </Grid>
        <Grid item xs={1}></Grid>
        <Grid item xs={1}></Grid>
        <Grid item xs={10}>
          <Button color="primary" type="submit" fullWidth variant="contained">
            Update Comment
          </Button>
        </Grid>
        <Grid item xs={1}></Grid>
      </Grid>
    </form>
  );
}

// type checking for props
UpdateCommentComponet.propTypes = {
  commentId: PropTypes.string,
  Comment: PropTypes.string,
  updateComment: PropTypes.func,
};

// setting default props
UpdateCommentComponet.defaultProps = {
  commentId: "",
  Comment: "",
  updateComment: () => {},
};

export default UpdateCommentComponet;
