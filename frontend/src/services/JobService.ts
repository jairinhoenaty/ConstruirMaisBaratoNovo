import Api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";
import { IJob } from "../interfaces/IJob";

const saveJob = (data: IJob) => ApiPublica.post("/job", data);

const getJobs = (limit: number, offset: number) =>
  Api.get(`/jobs?limit=${limit}&offset=${offset}`);

export const JobService = {
  saveJob,
  getJobs,
};