const commentsReducer = function (state = { comments: [] }, action) {
  switch (action.type) {
    case "GET_ALL_COMMENTS":
      return { ...state, comments: action.payload };
    default:
      return state;
  }
};
export default commentsReducer;
