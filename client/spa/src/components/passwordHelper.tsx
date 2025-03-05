import React, { useEffect, useState } from "react";

interface ValidationResult {
  rule: string;
  message: string;
  value?: number;
  valid: boolean;
}

const lengthRules = [
  { rule: "length", message: "A senha deve ter pelo menos 8 caracteres", value: 8 },
  { rule: "uppercase", message: "A senha deve conter pelo menos uma letra maiúscula" },
  { rule: "lowercase", message: "A senha deve conter pelo menos uma letra minúscula" },
  { rule: "number", message: "A senha deve conter pelo menos um número" },
  { rule: "special", message: "A senha deve conter pelo menos um caractere especial" }
];

function validatePassword(password: string): ValidationResult[] {
  return lengthRules.map((rule) => {
    switch (rule.rule) {
      case "length":
        return { ...rule, valid: password.length >= (rule.value ?? 0) };
      case "uppercase":
        return { ...rule, valid: /[A-Z]/.test(password) };
      case "lowercase":
        return { ...rule, valid: /[a-z]/.test(password) };
      case "number":
        return { ...rule, valid: /[0-9]/.test(password) };
      case "special":
        return { ...rule, valid: /[^A-Za-z0-9]/.test(password) };
      default:
        return { ...rule, valid: false };
    }
  });
}

const PasswordHelper = ({ password, focus }: { password: string; focus: boolean }) => {
  const [validationResults, setValidationResults] = useState<ValidationResult[]>([]);

  useEffect(() => {
    setValidationResults(validatePassword(password));
  }, [password]);

  return (
    <>
      {focus && !validationResults.every((result) => result.valid) && (
        <div className="flex flex-col gap-1 border p-4 rounded-lg bg-white absolute left-0" style={{
          top: "calc(100% + 0.5rem)", 
          }}>
          {validationResults.map((result, index) => (
            !result.valid && (
              <div key={index} className="flex items-center gap-2">
                <p className="flex gap-2 text-sm">
                  <span className="text-red-500 pt-[2px]">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="14"
                      height="14"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      strokeWidth="2"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      className="lucide lucide-x"
                    >
                      <path d="M18 6 6 18" />
                      <path d="m6 6 12 12" />
                    </svg>
                  </span>
                  {result.message}
                </p>
              </div>
            )
          ))}
        </div>
      )}
    </>
  );
};

export default PasswordHelper;