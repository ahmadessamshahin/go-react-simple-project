import { Fragment, PropsWithChildren } from 'react';
import { Avatar, Divider, ListItem, ListItemText, ListItemAvatar } from '@material-ui/core';

type ViewFieldProps = {
    title: string;
    content: string;
    divider: boolean
}

export default function ViewField(props: PropsWithChildren<ViewFieldProps>) {
    return (
        <Fragment>
            <ListItem>
                <ListItemAvatar>
                    <Avatar >
                        {props.children}
                    </Avatar>
                </ListItemAvatar>
                <ListItemText primary={props.title} secondary={props.content} />
            </ListItem>
            {props.divider && <Divider variant="inset" component="li" />}
        </Fragment>
    );
}
