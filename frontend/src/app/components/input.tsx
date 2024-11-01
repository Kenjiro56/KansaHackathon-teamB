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
    <div>

    <label
      htmlFor={id} className="text-gray-700">
        {label}
      </label>
      <input
        id={id}
        type={type}
        value={value}
        onChange={onChange}
        className="
        block
        rounded-md

        px-6
        pt-7
        pb-3
        w-full
        text-md
        text-gray-900
        appearance-none
        focus:outline-none
        focus:ring-0
        peer
      "
      placeholder=" "
      />
    </div>

  );
}


export default LoginInput;
