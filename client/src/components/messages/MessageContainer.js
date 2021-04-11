import React from "react";
import { Button, Card, CardContent, Grid, Typography } from "@material-ui/core";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { useDispatch } from "react-redux";
import axios from "axios";
import { getAllComments, getAllMessage } from "../../actions/index";
import CommentsContainer from "../comments/CommentsContainer.js";
import UpdateMessageComponet from "./UpdateMessageComponent";

function MessageContainer(props) {
  const dispatch = useDispatch();
  const [updateMesasgeStatus, setUpdateMessage] = React.useState(false);

  const getComments = () => {
    props.getComments(props._id);
  };

  const updateMessage = (message) => {
    axios
      .put("api/updateMessage", message, {
        params: {
          messageId: props._id,
        },
        headers: {
          "Content-Type": "application/json",
        },
      })
      .then(() => {
        dispatch(getAllMessage());
        setUpdateMessage(false);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const deleteMessage = () => {
    props.deleteMessage(props._id);
  };

  const createComment = (commentData) => {
    commentData.messageId = props._id;
    axios
      .post("api/createComment", commentData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then(() => {
        dispatch(getAllComments(props._id));
        const createComment = document.getElementById("create-comment");
        createComment.reset();
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const deleteComment = (commentId) => {
    axios
      .delete("api/deleteComment", {
        params: {
          commentId,
        },
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then(() => {
        dispatch(getAllComments(props._id));
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const updateComment = (commentData, commentId) => {
    axios
      .put("api/updateComment", commentData, {
        params: {
          commentId: commentId,
        },
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then(() => {
        dispatch(getAllComments(props._id));
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <React.Fragment>
      <Grid item xs={2}></Grid>
      <Grid item xs={8}>
        <Card>
          <CardContent>
            <Grid container spacing={0}>
              <Grid item xs={10}>
                {!updateMesasgeStatus ? (
                  <Typography variant="body1">{props.message}</Typography>
                ) : (
                  <UpdateMessageComponet
                    updateMessage={updateMessage}
                    Message={props.message}
                  />
                )}
              </Grid>

              {props.currentUserToken === props.createdBy && (
                <React.Fragment>
                  <Grid item xs={1}>
                    <Button
                      varaint="contained"
                      color="secondary"
                      size="small"
                      onClick={deleteMessage}
                    >
                      Delete
                    </Button>
                  </Grid>
                  <Grid item xs={1}>
                    <Button
                      varaint="contained"
                      color="primary"
                      size="small"
                      onClick={() => {
                        setUpdateMessage(!updateMesasgeStatus);
                      }}
                    >
                      Update
                    </Button>
                  </Grid>
                </React.Fragment>
              )}
              <Grid item xs={8}></Grid>
              <Grid item xs={4}>
                <Typography variant="caption">{props.createdBy}</Typography>
              </Grid>
              <Grid item xs={8}></Grid>
              <Grid item xs={4}>
                <Typography variant="caption">
                  {props.createdDate.toString()}
                </Typography>
              </Grid>
              <Grid item xs={12}>
                <Button variant="text" size="small" onClick={getComments}>
                  {props.viewComments ? "Hide Comments" : "View Comments"}
                </Button>
              </Grid>
              {props.viewComments && (
                <CommentsContainer
                  comments={props.comments}
                  currentUserToken={props.currentUserToken}
                  deleteComment={deleteComment}
                  createComment={createComment}
                  updateComment={updateComment}
                />
              )}
            </Grid>
          </CardContent>
        </Card>
      </Grid>
      <Grid item xs={2}></Grid>
    </React.Fragment>
  );
}

function mapStateToProps(state) {
  return {
    comments: state.comments.comments,
    currentUserToken: state.auth.currentUserToken,
  };
}

// type checking for props
MessageContainer.propTypes = {
  _id: PropTypes.string,
  message: PropTypes.string,
  createdBy: PropTypes.string,
  createdDate: PropTypes.objectOf(Date),
  comments: PropTypes.arrayOf(Object),
  viewComments: PropTypes.bool,
  currentUserToken: PropTypes.string,
  getComments: PropTypes.func,
  deleteMessage: PropTypes.func,
};

// setting default props
MessageContainer.defaultProps = {
  _id: "",
  message: "",
  createdBy: "",
  createdDate: null,
  comments: [],
  viewComments: false,
  currentUserToken: null,
  getComments: () => {},
  deleteMessage: () => {},
};
export default connect(mapStateToProps)(MessageContainer);
