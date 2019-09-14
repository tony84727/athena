import React from "react"
import Grid from "@material-ui/core/Grid"
import Button from "@material-ui/core/Button"
import TextField from "@material-ui/core/TextField"
import Typography from "@material-ui/core/Typography"

export default function AutoResponder() {
    return (
        <Grid container direction="row">
            <Typography variant="h1">
                Auto Responder
            </Typography>
            <Grid container>
                <Grid container spacing={2}>
                    <Grid item>
                        <TextField 
                            placeholder="Pattern"
                        />
                    </Grid>
                    <Grid item>
                        <TextField 
                            placeholder="Response"
                        />
                    </Grid>
                    <Grid item>
                        <Button>Add</Button>
                    </Grid>
                </Grid>
            </Grid>
        </Grid>
    )
}