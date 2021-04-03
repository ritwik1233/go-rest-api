const messageReducers = function (state = { messages: [] }, action) {
  switch (action.type) {
    case "GET_ALL_MESSAGE":
      return { ...state, messages: action.payload };
    default:
      return state;
  }
};
export default messageReducers;
