interface InputProps {
  id: string;
  onChange: any;
  value: string;
  label: string;
  type?: string;
}

const LoginInput: React.FC<InputProps> = ({
  id,
  onChange,
  value,
  label,
  type,
}) => {
  return (
    <div className="input-container">
      <input
        id={id}
        type={type}
        value={value}
        onChange={onChange}
      />
    </div>
    <label
    hetmlFor={id}>
      {label}
    </label>
  );
}


export default LoginInput;
