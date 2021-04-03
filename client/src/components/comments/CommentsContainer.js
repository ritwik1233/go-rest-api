import React from "react";
import { Grid, Typography, Button } from "@material-ui/core";
import PropTypes from "prop-types";
import CreateCommentComponent from "./CreateCommentComponent";
import UpdateCommentComponet from "./UpdateCommentComponent";

function CommentsContainer(props) {
  const [updateCommentId, setUpdateCommentStatus] = React.useState(null);

  const createComment = (commentData) => {
    props.createComment(commentData);
  };

  const updateComment = (commentData, commentId) => {
    props.updateComment(commentData, commentId);
    setUpdateCommentStatus(null);
  };

  return (
    <React.Fragment>
      <Grid item xs={12} style={{ margin: "0px", padding: "0px" }}>
        {props.comments.length > 0 ? (
          props.comments.map((comment) => {
            return (
              <Grid
                container
                spacing={3}
                key={comment.ID}
                style={{ margin: "0px", padding: "0px" }}
              >
                <Grid item xs={12} style={{ margin: "0px", padding: "0px" }}>
                  <Grid
                    container
                    spacing={3}
                    style={{ margin: "0px", padding: "0px" }}
                  >
                    <Grid
                      item
                      xs={10}
                      style={{ margin: "0px", padding: "0px" }}
                    >
                      {updateCommentId !== comment.ID ? (
                        <Typography
                          variant="caption"
                          style={{ margin: "0px", padding: "0px" }}
                        >
                          {comment.Comment}
                        </Typography>
                      ) : (
                        <UpdateCommentComponet
                          commentId={comment.ID}
                          Comment={comment.Comment}
                          updateComment={updateComment}
                        />
                      )}
                    </Grid>
                    {props.userEmail === comment.CreatedBy && (
                      <React.Fragment>
                        <Grid
                          item
                          xs={1}
                          style={{ margin: "0px", padding: "0px" }}
                        >
                          <Button
                            varaint="contained"
                            color="secondary"
                            size="small"
                            onClick={() => {
                              props.deleteComment(comment.ID);
                            }}
                          >
                            Delete
                          </Button>
                        </Grid>
                        <Grid
                          item
                          xs={1}
                          style={{ margin: "0px", padding: "0px" }}
                        >
                          <Button
                            varaint="contained"
                            color="primary"
                            size="small"
                            onClick={() => {
                              if (updateCommentId === comment.ID) {
                                setUpdateCommentStatus(null);
                                return;
                              }
                              setUpdateCommentStatus(comment.ID);
                            }}
                          >
                            Update
                          </Button>
                        </Grid>
                      </React.Fragment>
                    )}
                    {updateCommentId !== comment.ID && (
                      <React.Fragment>
                        <Grid
                          item
                          xs={8}
                          style={{ margin: "0px", padding: "0px" }}
                        ></Grid>
                        <Grid
                          item
                          xs={4}
                          style={{ margin: "0px", padding: "0px" }}
                        >
                          <Typography
                            variant="caption"
                            style={{ margin: "0px", padding: "0px" }}
                          >
                            {comment.CreatedBy}
                          </Typography>
                        </Grid>
                        <Grid
                          item
                          xs={8}
                          style={{ margin: "0px", padding: "0px" }}
                        ></Grid>
                        <Grid
                          item
                          xs={4}
                          style={{ margin: "0px", padding: "0px" }}
                        >
                          <Typography
                            variant="caption"
                            style={{ margin: "0px", padding: "0px" }}
                          >
                            {comment.CreatedDate.toString()}
                          </Typography>
                        </Grid>
                      </React.Fragment>
                    )}
                  </Grid>
                </Grid>
              </Grid>
            );
          })
        ) : (
          <Typography variant="body1">No Comments..</Typography>
        )}
      </Grid>
      <Grid item xs={12}>
        &nbsp;
      </Grid>
      {props.userEmail && (
        <CreateCommentComponent createComment={createComment} />
      )}
    </React.Fragment>
  );
}

// type checking for props
CommentsContainer.propTypes = {
  comments: PropTypes.arrayOf(Object),
  userEmail: PropTypes.string,
  deleteComment: PropTypes.func,
  createComment: PropTypes.func,
  updateComment: PropTypes.func,
};

// setting default props
CommentsContainer.defaultProps = {
  comments: [],
  userEmail: null,
  deleteComment: () => {},
  createComment: () => {},
  updateComment: () => {},
};
export default CommentsContainer;
