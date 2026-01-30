import React, { useState } from 'react';

interface FormData {
  content: string;
  rule_id: string;
  rule_name: string;
  deep: string;
  search_ql: string;
}

interface FormComponentProps {
  onSubmit: (data: FormData) => void;
  loading: boolean;
}

const FormComponent: React.FC<FormComponentProps> = ({ onSubmit, loading }) => {
  const [formData, setFormData] = useState<FormData>({
    content: '',
    rule_id: '',
    rule_name: '',
    deep: 'true',
    search_ql: '*'
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(formData);
  };

  return (
    <div className="form-container">
      <h2>Search Form</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="content">Content:</label>
          <textarea
            id="content"
            name="content"
            value={formData.content}
            onChange={handleChange}
            required
            placeholder="Enter content here..."
          />
        </div>

        <div className="form-group">
          <label htmlFor="rule_id">Rule ID:</label>
          <input
            type="text"
            id="rule_id"
            name="rule_id"
            value={formData.rule_id}
            onChange={handleChange}
            required
            placeholder="Enter rule ID"
          />
        </div>

        <div className="form-group">
          <label htmlFor="rule_name">Rule Name:</label>
          <input
            type="text"
            id="rule_name"
            name="rule_name"
            value={formData.rule_name}
            onChange={handleChange}
            required
            placeholder="Enter rule name"
          />
        </div>

        <div className="form-group">
          <label htmlFor="deep">Deep:</label>
          <select
            id="deep"
            name="deep"
            value={formData.deep}
            onChange={handleChange}
            required
          >
            <option value="true">true</option>
            <option value="false">false</option>
          </select>
        </div>

        <div className="form-group">
          <label htmlFor="search_ql">Search QL:</label>
          <textarea
            id="search_ql"
            name="search_ql"
            value={formData.search_ql}
            onChange={handleChange}
            required
            placeholder="Enter search query language"
          />
        </div>

        <button type="submit" className="submit-button" disabled={loading}>
          {loading ? 'Submitting...' : 'Submit'}
        </button>
      </form>
    </div>
  );
};

export default FormComponent;
export type { FormData };
