import axios from "axios";

export const uploadFileToServer = async (formData: FormData): Promise<string> => {
    try {
        const response = await axios.post("http://localhost:8000/files", formData, {
            headers: {
                "Content-Type": "multipart/form-data",
            },
        });

        console.log("Response : ", response)
        console.log("File Url : ", response.data.fileUrl)

        return response.data.fileUrl;
    } catch (error) {
        console.error("Error uploading file:", error);
        if (error instanceof axios.AxiosError) {
            throw new Error("File upload failed: " + error.response?.data);
        }else{
                    throw new Error("File upload failed. Please try again.");
        }
    }
};
