import axios from 'axios';

const API_BASE = 'http://localhost:8080/api';

export const getJobs = () => axios.get<Job[]>(`${API_BASE}/jobs`);
export const submitApplication = (app: Application) => axios.post(`${API_BASE}/applications`, app);
