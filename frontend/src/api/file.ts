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

        if (response.status === 200) {
            return response.data.fileUrl;
        } else {
            throw new Error("File upload failed. Please try again.");
        }
    } catch (error) {
        console.error("Error uploading file:", error);
        throw new Error("File upload failed. Please try again.");
    }
};
