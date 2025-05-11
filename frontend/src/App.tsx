import { HardDriveDownload } from "lucide-react"
import { useEffect, useRef, useState, type ChangeEvent } from "react"

const App = () => {
  const fileInput = useRef<HTMLInputElement | null>(null);
  const [file, setFile] = useState<File | undefined>();
  const [error, setError] = useState("")

  const handleUploadClick = () => {
    if (!fileInput.current) return;

    fileInput.current.click();
  }

  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    if (!e.target.files) return;
    setError("")

    const newFile = e.target.files[0]
    if (newFile.size > 1058336) {
      setError("Size limit exceeded: File should be smaller than 1 MB")
    }
    setFile(newFile);
  };

  useEffect(() => {
    console.log("File changed: ", file)
    if (!file) return;

    const data = new FormData();
    data.append("file", file);
  }, [file])

  // useEffect(() => {
  //   console.log("FileInput: ", fileInput.current)
  // }, [fileInput.current])


  return (
    <main className="h-screen w-screen bg-slate-950 flex items-center justify-center">
      <div className="bg-slate-700 min-h-1/2 w-1/2 min-w-xs rounded-2xl">
        <div className="w-full my-4 flex items-center justify-center gap-4">
          <HardDriveDownload className="text-blue-500" size={32} />
          <h1 className="w-fit text-3xl font-bold text-slate-300">QuickDrop</h1>
        </div>
        <h2 className="max-w-4/5 m-auto text-slate-300 font-light text-center border-2 border-blue-500 p-2 rounded-md">
          A minimal full-stack app to upload a file and instantly generate a unique link to download it.
        </h2>

        <p className="text-slate-100 text-center text-xs mt-4">Upload file upto 1 MB</p>
        <div className="w-fit my-2 mx-auto">
          <button
            className="bg-blue-500 px-4 py-2 rounded-lg font-semibold text-slate-300 hover:bg-blue-600"
            onClick={handleUploadClick}
          >Upload</button>
          <input type="file" className="hidden" ref={fileInput} onChange={handleFileChange} />
        </div>

        <p className="text-red-500 text-center text-xs my-2">{error}</p>
      </div>
    </main>
  )
}

export default App