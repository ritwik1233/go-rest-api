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
          payload: res.data.result,
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
          payload: res.data.result,
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

export function loginUser(userDetails) {
  return function (dispatch) {
    axios({
      method: "post",
      url: "/api/login",
      data: userDetails,
      headers: { "Content-Type": "multipart/form-data" },
    }).then((res) => {
      const access_token = `${res.data.auth}`;
      dispatch({
        type: "LOGIN_USER",
        payload: access_token,
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
