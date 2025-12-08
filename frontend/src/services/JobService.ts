import api from "../providers/Api";
import ApiPublica from "../providers/ApiPublica";

export const JobService = {

  // PÚBLICO — cadastrar vaga
  create(data: any) {
    return ApiPublica.post("/jobs", data);
  },

  // PRIVADO — ADMIN
getAll(limit = 20, offset = 0) {
  return api.get(`/jobs?limit=${limit}&offset=${offset}`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`
    }
  });
},



approve(id: number, approved: boolean) {
  return api.put(`/jobs/approve/${id}`, { approved }, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "User-Type": "admin"
    }
  });
},



  // PRIVADO — PREMIUM
 getApproved() {
    return ApiPublica.get("/jobs/approved");
  },



  delete(id: number) {
    return api.delete(`/jobs/${id}`);
  }
  
  }

