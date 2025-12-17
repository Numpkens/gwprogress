import { useState, ChangeEvent } from 'react'
import './App.css'

function App() {
  // 1. Memory: This stores the file the user picks. 
  // It starts as 'null' because no file is selected yet.
  const [file, setFile] = useState<File | null>(null);

  // 2. Selection: This runs when you pick a file from your computer.
  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      setFile(e.target.files[0]);
    }
  };

  // 3. Action: This will send the file to your Go backend later.
  const handleUpload = async () => {
    if (!file) {
      alert("Please select a GW2 log file first!");
      return;
    }

    console.log("Preparing to send to Go backend:", file.name);
    
    // We will write the 'fetch' code here once our Go server is ready!
    alert(`Ready to upload: ${file.name}`);
  };

  return (
    <div className="container">
      <h1>GW Progress</h1>
      <p>Upload your ArcDPS (.evtc) logs to analyze performance.</p>

      <div className="card">
        {/* The standard HTML file input */}
        <input 
          type="file" 
          accept=".evtc,.zip" 
          onChange={handleFileChange} 
        />

        <p>Selected: {file ? file.name : "None"}</p>

        <button onClick={handleUpload} disabled={!file}>
          Upload Log
        </button>
      </div>
    </div>
  )
}

export default App
