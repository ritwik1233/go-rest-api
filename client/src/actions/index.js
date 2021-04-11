import axios from "axios";

export function getAllMessage() {
  return function (dispatch) {
    axios
      .get("/api/getAllMessage", {
        headers: {
          "Content-Type": "application/json",
        },
      })
      .then((res) => {
        dispatch({
          type: "GET_ALL_MESSAGE",
          payload: res.data || [],
        });
      })
      .catch((err) => {
        console.error(err);
        dispatch({
          type: "GET_ALL_MESSAGE",
          payload: [],
        });
      });
  };
}

export function getAllComments(messageId) {
  return function (dispatch) {
    axios
      .get("/api/getComments", {
        headers: {
          "Content-Type": "application/json",
        },
        params: {
          messageId,
        },
      })
      .then((res) => {
        dispatch({
          type: "GET_ALL_COMMENTS",
          payload: res.data || [],
        });
      })
      .catch((err) => {
        console.error(err);
        dispatch({
          type: "GET_ALL_COMMENTS",
          payload: [],
        });
      });
  };
}

export function getCurrentUser() {
  return function (dispatch) {
    axios
      .get("/api/getuser", {
        headers: { "Content-Type": "application/json" },
      })
      .then((res) => {
        dispatch({
          type: "GET_USER",
          payload: res.data.email,
        });
      })
      .catch(() => {
        dispatch({
          type: "LOGOUT_USER",
        });
      });
  };
}

export function logoutUser(currentUserToken) {
  return function (dispatch) {
    axios
      .get("/api/logout", {
        headers: {
          "Content-Type": "application/json",
          Authorization: currentUserToken,
        },
      })
      .then((res) => {
        console.log(res.data.result);
        dispatch({
          type: "LOGOUT_USER",
        });
      })
      .catch(() => {
        dispatch({
          type: "LOGOUT_USER",
        });
      });
  };
}
