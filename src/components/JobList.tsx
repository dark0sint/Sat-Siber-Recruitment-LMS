import React, { useEffect, useState } from 'react';
import { Job } from '../types';
import { getJobs } from '../services/api';

const JobList: React.FC = () => {
  const [jobs, setJobs] = useState<Job[]>([]);

  useEffect(() => {
    getJobs().then(res => setJobs(res.data));
  }, []);

  return (
    <div>
      <h2>Job Listings - Sat Siber Recruitment LMS</h2>
      {jobs.map(job => (
        <div key={job.id}>
          <h3>{job.title}</h3>
          <p>{job.description}</p>
        </div>
      ))}
    </div>
  );
};

export default JobList;
