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
  const [userEmail, setUserEmail] = React.useState(null);
  const [updateMesasgeStatus, setUpdateMessage] = React.useState(false);
  React.useEffect(() => {
    axios
      .get("/api/getCurrentUser", {
        headers: {
          "Content-Type": "application/json",
          Authorization: props.currentUserToken,
        },
      })
      .then((res) => {
        setUpdateMessage(false);
        setUserEmail(res.data.result);
      })
      .catch(() => {
        setUserEmail(null);
      });
  }, [props.currentUserToken]);

  const getComments = () => {
    props.getComments(props.ID);
  };

  const updateMessage = (message) => {
    axios({
      method: "put",
      url: "api/updateMessage",
      data: message,
      params: {
        messageId: props.ID,
      },
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: props.currentUserToken,
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
    props.deleteMessage(props.ID);
  };

  const createComment = (commentData) => {
    commentData.append("messageId", props.ID);
    axios({
      method: "post",
      url: "api/createComment",
      data: commentData,
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: props.currentUserToken,
      },
    })
      .then(() => {
        dispatch(getAllComments(props.ID));
        const createComment = document.getElementById("create-comment");
        createComment.reset();
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const deleteComment = (commentId) => {
    axios({
      method: "delete",
      url: "api/deleteComment",
      params: {
        commentId,
      },
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: props.currentUserToken,
      },
    })
      .then(() => {
        dispatch(getAllComments(props.ID));
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const updateComment = (commentData, commentId) => {
    axios({
      method: "put",
      url: "api/updateComment",
      data: commentData,
      params: {
        commentId: commentId,
      },
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: props.currentUserToken,
      },
    })
      .then(() => {
        dispatch(getAllComments(props.ID));
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
                  <Typography variant="body1">{props.Message}</Typography>
                ) : (
                  <UpdateMessageComponet
                    updateMessage={updateMessage}
                    Message={props.Message}
                  />
                )}
              </Grid>

              {userEmail === props.CreatedBy && (
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
                <Typography variant="caption">{props.CreatedBy}</Typography>
              </Grid>
              <Grid item xs={8}></Grid>
              <Grid item xs={4}>
                <Typography variant="caption">
                  {props.CreatedDate.toString()}
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
                  userEmail={userEmail}
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
  ID: PropTypes.string,
  Message: PropTypes.string,
  CreatedBy: PropTypes.string,
  CreatedDate: PropTypes.objectOf(Date),
  comments: PropTypes.arrayOf(Object),
  viewComments: PropTypes.bool,
  currentUserToken: PropTypes.string,
  getComments: PropTypes.func,
  deleteMessage: PropTypes.func,
};

// setting default props
MessageContainer.defaultProps = {
  ID: "",
  Message: "",
  CreatedBy: "",
  CreatedDate: null,
  comments: [],
  viewComments: false,
  currentUserToken: null,
  getComments: () => {},
  deleteMessage: () => {},
};
export default connect(mapStateToProps)(MessageContainer);
