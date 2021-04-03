import React from "react";
import { Grid, Paper } from "@material-ui/core";
import { connect } from "react-redux";
import PropTypes from "prop-types";
import { useDispatch } from "react-redux";
import axios from "axios";

import { getAllMessage, getAllComments } from "../actions/index";
import MessageContainer from "../components/messages/MessageContainer";
import CreateMessageComponent from "../components/messages/CreateMessageComponent";

function HomePage(props) {
  const dispatch = useDispatch();
  const [messageId, setMessageId] = React.useState(null);

  React.useEffect(() => {
    dispatch(getAllMessage());
  }, []);

  const getComments = (_id) => {
    if (messageId === _id) {
      setMessageId(null);
      dispatch(getAllComments(""));
    } else {
      setMessageId(_id);
      dispatch(getAllComments(_id));
    }
  };

  const handleMessage = (message) => {
    axios({
      method: "post",
      url: "api/createMessage",
      data: message,
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: props.currentUserToken,
      },
    })
      .then(() => {
        dispatch(getAllMessage());
        const createMessage = document.getElementById("create-message");
        createMessage.reset();
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const deleteMessage = (messageId) => {
    axios({
      method: "delete",
      url: "api/deleteMessage",
      params: {
        messageId,
      },
      headers: {
        "Content-Type": "multipart/form-data",
        Authorization: props.currentUserToken,
      },
    })
      .then(() => {
        dispatch(getAllMessage());
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const messageComponent = props.messages.map((message) => {
    const formattedDate = new Date(message.CreatedDate);
    const viewComments = message.ID === messageId;
    return (
      <MessageContainer
        key={message.ID}
        ID={message.ID}
        Message={message.Message}
        CreatedBy={message.CreatedBy}
        CreatedDate={formattedDate}
        getComments={getComments}
        deleteMessage={deleteMessage}
        viewComments={viewComments}
      />
    );
  });

  return (
    <Grid container spacing={3}>
      <React.Fragment>
        <Grid item xs={2}></Grid>
        <Grid item xs={8}>
          {props.currentUserToken && (
            <Paper>
              <CreateMessageComponent handleMessage={handleMessage} />
            </Paper>
          )}
        </Grid>
        <Grid item xs={2}></Grid>
        {messageComponent}
      </React.Fragment>
    </Grid>
  );
}

function mapStateToProps(state) {
  return {
    messages: state.message.messages,
    currentUserToken: state.auth.currentUserToken,
  };
}

// type checking for props
HomePage.propTypes = {
  messages: PropTypes.arrayOf(Object),
  currentUserToken: PropTypes.string,
};

// setting default props
HomePage.defaultProps = {
  messages: [],
  currentUserToken: null,
};
export default connect(mapStateToProps)(HomePage);
