import React from 'react';
import { useHistory } from 'react-router-dom';
import {
    Button,
    Menu,
    MenuItem
} from '@material-ui/core';
import MoreVertIcon from '@material-ui/icons/MoreVert';

type DropDownProps = {
    phoneNumberId: number;
}

export default function DropDown(props: DropDownProps) {
    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
    const history = useHistory();

    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = (path: string) => {
        history.push(path);
    };

    return (
        <div>
            <Button aria-controls="simple-menu" aria-haspopup="true" onClick={handleClick}>
                <MoreVertIcon />
            </Button>
            <Menu
                id="simple-menu"
                anchorEl={anchorEl}
                keepMounted
                open={Boolean(anchorEl)}
                onClose={handleClose}
            >
                <MenuItem onClick={() => handleClose(`/${props.phoneNumberId}`)}>View</MenuItem>
                <MenuItem onClick={() => handleClose('')}>Edit</MenuItem>
                <MenuItem onClick={() => handleClose('')}>Delete</MenuItem>
            </Menu>
        </div>
    );
}