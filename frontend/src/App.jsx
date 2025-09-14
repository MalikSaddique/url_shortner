import { useState } from 'react'
import './App.css'

function App() {
  const [long_url, setUrl] = useState('')
  const [expirationDays, setExpirationDays] = useState(30)
  const [response, setResponse] = useState(null)
  const backendURL = import.meta.env.VITE_BACKEND_URL

  const handleChange = (e) => {
    setUrl(e.target.value)
  }

  const handleExpirationChange = (e) => {
    setExpirationDays(parseInt(e.target.value) || 30)
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      const requestBody = { long_url }
      
      if (expirationDays !== 30) {
        const expirationDate = new Date()
        expirationDate.setDate(expirationDate.getDate() + expirationDays)
        requestBody.expire_at = expirationDate.toISOString()
      }

      const res = await fetch(`${backendURL}/createurl`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestBody)
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

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    })
  }

  const getTimeRemaining = (expireAt) => {
    const now = new Date()
    const expire = new Date(expireAt)
    const diff = expire.getTime() - now.getTime()
    
    if (diff <= 0) return 'Expired'
    
    const days = Math.floor(diff / (1000 * 60 * 60 * 24))
    const hours = Math.floor((diff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
    const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60))
    
    if (days > 0) return `${days} days, ${hours}h remaining`
    if (hours > 0) return `${hours}h ${minutes}m remaining`
    return `${minutes} minutes remaining`
  }

  return (
    <>
      <div className="container">
        <h1>URL Shortener</h1>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <input
              type="text"
              value={long_url}
              onChange={handleChange}
              placeholder="Enter URL to shorten"
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="expiration">Expires in:</label>
            <select 
              id="expiration" 
              value={expirationDays} 
              onChange={handleExpirationChange}
            >
              <option value={1}>1 day</option>
              <option value={7}>7 days</option>
              <option value={30}>30 days</option>
              <option value={90}>90 days</option>
              <option value={365}>1 year</option>
            </select>
          </div>
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
                  <div className="url-row">
                    <label>Expires at:</label>
                    <div className="expiration-info">
                      <span className="expiration-date">
                        {formatDate(response.expire_at)}
                      </span>
                      <span className="time-remaining">
                        {getTimeRemaining(response.expire_at)}
                      </span>
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