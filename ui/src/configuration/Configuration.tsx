import * as React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import { Title } from 'react-admin';
import { makeStyles } from '@mui/styles';
import { ToggleThemeButton } from 'react-admin';
import { darkTheme, lightTheme } from './themes';


const useStyles = makeStyles({
    label: { width: '10em', display: 'inline-block' },
    button: { margin: '1em' },
});

const Configuration = () => {
    const classes = useStyles();
    return (
        <Card>
            <Title title='Configuration' />
            <CardContent>
                <div className={classes.label}>
                    Theme
                </div>
                <ToggleThemeButton lightTheme={lightTheme} darkTheme={darkTheme} />
            </CardContent>
        </Card>
    );
};

export default Configuration;
