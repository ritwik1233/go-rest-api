import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { createStore, applyMiddleware } from "redux";
import reduxThunk from "redux-thunk";
import reducers from "./reducers";
import { persistStore, persistReducer } from "redux-persist";
import { PersistGate } from "redux-persist/integration/react";
import storage from "redux-persist/lib/storage";
import { Grid } from "@material-ui/core";
import { BrowserRouter, Route, Switch } from "react-router-dom";

import Header from "./components/Header.js";
import HomePage from "./pages/HomePage.js";
import NotFound from "./pages/NotFound.js";

const persistConfig = {
  key: "root",
  storage,
};
const persistedReducer = persistReducer(persistConfig, reducers);
let store = createStore(persistedReducer, {}, applyMiddleware(reduxThunk));
let persistor = persistStore(store);

ReactDOM.render(
  <Provider store={store}>
    <PersistGate loading={null} persistor={persistor}>
      <BrowserRouter>
        <Grid container spacing={3}>
          <Grid item xs={12}>
            <Header />
          </Grid>
          <Grid item xs={12}>
            <Switch>
              <Route exact path="/" component={HomePage} />
              <Route component={NotFound} />
            </Switch>
          </Grid>
        </Grid>
      </BrowserRouter>
    </PersistGate>
  </Provider>,
  document.getElementById("root")
);
