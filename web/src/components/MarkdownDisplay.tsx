import React from 'react';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';

interface MarkdownDisplayProps {
  content: string;
  loading?: boolean;
}

const MarkdownDisplay: React.FC<MarkdownDisplayProps> = ({ content, loading }) => {
  if (loading) {
    return <div className="loading">Loading...</div>;
  }

  if (!content) {
    return null;
  }

  return (
    <div className="result-container">
      <h2>Result</h2>
      <div className="markdown-content">
        <ReactMarkdown remarkPlugins={[remarkGfm]}>
          {content}
        </ReactMarkdown>
      </div>
    </div>
  );
};

export default MarkdownDisplay;
