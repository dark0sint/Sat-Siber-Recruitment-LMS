export interface Job {
  id: number;
  title: string;
  category: string;
  location: string;
  department: string;
  description: string;
}

export interface Application {
  job_id: number;
  name: string;
  email: string;
  resume: string;
}
