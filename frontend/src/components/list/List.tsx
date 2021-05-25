import React, { useEffect } from 'react';
import {
    Paper,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TablePagination,
    TableRow
} from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import columns from './Columns';
import DropDown from './Dropdown';
import { useDispatch, useSelector } from 'react-redux'
import { getPhoneNumber } from '../../store/action'
import phoneNumbers from './PhoneNumbers'
const useStyles = makeStyles({
    root: {
        width: '100%',
    }
});

export default function List() {
    const classes = useStyles();
    const [page, setPage] = React.useState(0);
    const [rowsPerPage, setRowsPerPage] = React.useState(10);

    const handleChangePage = (event: unknown, newPage: number) => {
        setPage(newPage);
    };
    const dispatch = useDispatch()
    const phoneList = useSelector((state:any) => state.phoneNumberList)
    console.log(phoneList);
    
    const { loading, error, data } = phoneList

    useEffect(() => {
        dispatch(getPhoneNumber())
    }, [dispatch])

    const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
        setRowsPerPage(+event.target.value);
        setPage(0);
    };

    return (
        <Paper className={classes.root}>
            <TableContainer>
                <Table stickyHeader aria-label="sticky table">
                    <TableHead>
                        <TableRow>
                            {columns.map((column) => (
                                <TableCell
                                    key={column.id}
                                    style={{ minWidth: column.minWidth }}
                                >
                                    {column.label}
                                </TableCell>
                            ))}
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {phoneNumbers.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map((phoneNumber) => {
                            return (
                                <TableRow hover role="checkbox" tabIndex={-1} key={phoneNumber.id}>
                                    {columns.map((column) => {
                                        return (
                                            <TableCell key={column.id}>
                                                {column.id !== 'action' ? phoneNumber[column.id] : <DropDown phoneNumberId={phoneNumber.id} />}
                                            </TableCell>
                                        );
                                    })}
                                </TableRow>
                            );
                        })}
                    </TableBody>
                </Table>
            </TableContainer>
            <TablePagination
                rowsPerPageOptions={[10, 25, 100]}
                component="div"
                count={phoneNumbers.length}
                rowsPerPage={rowsPerPage}
                page={page}
                onChangePage={handleChangePage}
                onChangeRowsPerPage={handleChangeRowsPerPage}
            />
        </Paper>
    );
}
