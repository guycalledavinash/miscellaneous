import {TextField} from '@mui/material';
export function SearchBar({value,onChange}:{value:string;onChange:(v:string)=>void}){return <TextField fullWidth margin="normal" label="Search players, clubs" value={value} onChange={e=>onChange(e.target.value)}/>}
