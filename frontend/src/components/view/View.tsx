import { List } from '@material-ui/core';
import { Theme, createStyles, makeStyles } from '@material-ui/core/styles';
import { Phone, Flag, Done } from '@material-ui/icons';
import ViewField from './ViewField';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            width: '100%',
            backgroundColor: theme.palette.background.paper,
        },
    }),
);

export default function View() {
    const classes = useStyles();

    const viewFields = [
        {
            title: 'Phone Number',
            content: '(20) 1091083167', // to be replaced..
            divider: true,
            icon: <Phone />
        },
        {
            title: 'Country',
            content: 'Egypt', // to be replaced..
            divider: true,
            icon: <Flag />
        },
        {
            title: 'Status',
            content: 'Valid', // to be replaced..
            divider: false,
            icon: <Done />
        }
    ]

    return (
        <List className={classes.root}>
            {viewFields.map((viewField) => {
                return (
                    <ViewField
                        key={viewField.title}
                        title={viewField.title}
                        content={viewField.content}
                        divider={viewField.divider}
                    >
                        {viewField.icon}
                    </ViewField>
                );
            })}
        </List>
    );
}