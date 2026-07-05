import axios from 'axios';
export const api=axios.create({baseURL:import.meta.env.VITE_API_URL??'/api'});
api.interceptors.request.use((config)=>{const token=localStorage.getItem('token'); if(token) config.headers.Authorization=`Bearer ${token}`; return config;});
export async function login(username:string,password:string){const {data}=await api.post('/auth/login',{username,password}); localStorage.setItem('token',data.access_token); return data;}
