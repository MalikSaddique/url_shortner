import { useState } from 'react'
import './App.css'

function App() {
  const [long_url, setUrl] = useState('')
  const [response, setResponse] = useState(null)
  const backendURL = import.meta.env.VITE_BACKEND_URL

  const handleChange = (e) => {
    setUrl(e.target.value)
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      const res = await fetch(`${backendURL}/createurl`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ long_url })
      })
      const data = await res.json()
      setResponse(data)
    } catch (error) {
      setResponse({ error: 'Failed to create URL' })
    }
  }

  const handleRedirect = (shortUrl) => {
    window.open(`${backendURL}/${shortUrl}`, '_blank')
  }

  return (
    <>
      <div className="container">
        <h1>URL Shortener</h1>
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            value={long_url}
            onChange={handleChange}
            placeholder="Enter URL to shorten"
            required
          />
          <button type="submit">Shorten URL</button>
        </form>
        
        {response && (
          <div className="response-container">
            {response.error ? (
              <div className="error">
                <h3>Error:</h3>
                <p>{response.error}</p>
              </div>
            ) : (
              <div className="success">
                <h3>Short URL Created Successfully!</h3>
                <div className="url-info">
                  <div className="url-row">
                    <label>Original URL:</label>
                    <a href={response.long_url} target="_blank" rel="noopener noreferrer">
                      {response.long_url}
                    </a>
                  </div>
                  <div className="url-row">
                    <label>Short URL:</label>
                    <div className="short-url-container">
                      <span className="short-url">{backendURL}/{response.short_url}</span>
                      <button 
                        className="copy-btn"
                        onClick={() => navigator.clipboard.writeText(`${backendURL}/${response.short_url}`)}
                        title="Copy to clipboard"
                      >
                        📋
                      </button>
                      <button 
                        className="test-btn"
                        onClick={() => handleRedirect(response.short_url)}
                        title="Test redirect"
                      >
                        🔗
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            )}
          </div>
        )}
      </div>
    </>
  )
}

export default App
