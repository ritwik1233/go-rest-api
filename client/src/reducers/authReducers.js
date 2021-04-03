const authReducer = function (state = { currentUserToken: null }, action) {
  switch (action.type) {
    case "LOGIN_USER":
      return { ...state, currentUserToken: action.payload };
    case "LOGOUT_USER":
      return { ...state, currentUserToken: null };
    default:
      return state;
  }
};
export default authReducer;
