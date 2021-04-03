import React from "react";
import { Grid, TextField, Button } from "@material-ui/core";
import PropTypes from "prop-types";

function CreateCommentComponent(props) {
  const createComment = (e) => {
    e.preventDefault();
    const commentData = new FormData();
    commentData.append("comment", e.target.comment.value);
    props.createComment(commentData);
  };

  return (
    <Grid item xs={12}>
      <form id="create-comment" onSubmit={createComment}>
        <Grid container spacing={3}>
          <Grid item xs={12}>
            <TextField
              name="comment"
              fullWidth
              multiline
              label="Add Comment"
              variant="outlined"
              required
            />
          </Grid>
          <Grid item xs={12}>
            <Button color="primary" type="submit" fullWidth variant="contained">
              Add Comment
            </Button>
          </Grid>
        </Grid>
      </form>
    </Grid>
  );
}

// type checking for props
CreateCommentComponent.propTypes = {
  createComment: PropTypes.func,
};

// setting default props
CreateCommentComponent.defaultProps = {
  createComment: () => {},
};

export default CreateCommentComponent;
