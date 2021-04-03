import React from "react";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";

function NotFound() {
  return (
    <Grid container spacing={3}>
      <Grid item xs={12}>
        <Typography variant="h4">Page Not Found</Typography>
      </Grid>
    </Grid>
  );
}

export default NotFound;
