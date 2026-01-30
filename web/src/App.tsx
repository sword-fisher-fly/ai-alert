import React, { useState } from 'react';
import FormComponent, { FormData } from './components/FormComponent';
import MarkdownDisplay from './components/MarkdownDisplay';
import './App.css';

function App() {
  const [result, setResult] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');

  const handleSubmit = async (data: FormData) => {
    setLoading(true);
    setError('');
    setResult('');

    try {
      const formBody = Object.keys(data)
        .map(key => encodeURIComponent(key) + '=' + encodeURIComponent(data[key as keyof FormData]))
        .join('&');

      const response = await fetch('http://localhost:8090/api/v1/ai/chat', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
          'Authorization': 'Bearer mock-token-12345',
        },
        body: formBody,
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const responseData = await response.json();

      // Extract content from response.data field
      let markdownContent = '';
      if (responseData.data) {
        const data = responseData.data;
        if (typeof data === 'string') {
          markdownContent = data;
        } else if (data.content) {
          markdownContent = data.content;
        } else if (data.message) {
          markdownContent = data.message;
        } else {
          markdownContent = JSON.stringify(data, null, 2);
        }
      } else if (typeof responseData === 'string') {
        markdownContent = responseData;
      } else if (responseData.content) {
        markdownContent = responseData.content;
      } else if (responseData.message) {
        markdownContent = responseData.message;
      } else {
        markdownContent = JSON.stringify(responseData, null, 2);
      }

      setResult(markdownContent);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
      console.error('API Error:', err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="App">
      <div className="container">
        <h1>AI诊断运维告警信息</h1>
        <div className="main-layout">
          <div className="left-panel">
            <FormComponent onSubmit={handleSubmit} loading={loading} />
            {error && <div className="error">Error: {error}</div>}
          </div>
          <div className="right-panel">
            <MarkdownDisplay content={result} loading={loading} />
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
