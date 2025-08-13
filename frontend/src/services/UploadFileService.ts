import axios from "axios";

/*const uploadImage = (data: any) =>
  ApiPublica.post("/image/upload", data);*/


const uploadImage = async (file:string) => {
    console.log("uploadImage", file);
  const baseURL = import.meta.env.VITE_BASE_URL_PUBLICA;
  const response = await axios({
    url: baseURL + "/upload/image", //your url
    method: "POST",
    responseType: "json", // important
    headers: {
      "Content-Type": "multipart/form-data",
    },
    data: {
      myFile: file // Assuming you have an input of type file       
    },
  })
  /*.then((response) => {
    // create file link in browser's memory
    console.log(response.headers);
    //const href = URL.createObjectURL(response.data);
    console.log("data", response.data);  
    return response.data; // return the file data

    // create "a" HTML element with href to file & click
  });*/
  return response
};
  

export const UploadFileService = {
  uploadImage,
};
