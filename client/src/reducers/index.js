import { combineReducers } from "redux";
import authReducers from "./authReducers";
import messageReducers from "./messageReducers";
import commentsReducer from "./commentReducers";

const CombineReducers = combineReducers({
  auth: authReducers,
  message: messageReducers,
  comments: commentsReducer,
});

export default CombineReducers;
