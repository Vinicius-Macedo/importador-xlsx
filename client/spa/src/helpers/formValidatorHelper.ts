interface ValidationResult {
  [key: string]: string;
}

export function validateFields(fields: {
  [key: string]: string;
}): ValidationResult {
  const errors: ValidationResult = {};

  for (const [field, value] of Object.entries(fields)) {
    let error = "";
    switch (field) {
      case "email":
        error = validateEmail(value);
        break;
      case "password":
        error = validatePassword(value);
        break;
      case "username":
        error = validateUsername(value);
        break;
      case "name":
      case "firstName":
      case "lastName":
        error = validateName(value);
        break;
      case "repeat_password":
        error = validateRepeatPassword(fields.password, value);
        break;
      default:
        break;
    }
    if (error) {
      errors[field] = error;
    }
  }

  return errors;
}

export function validateEmail(email: string): string {
  if (email === "") {
    return "O email é obrigatório";
  }

  if (email.length > 255) {
    return "O email é muito longo";
  }

  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!emailRegex.test(email)) {
    return "O email é inválido";
  }

  return "";
}

export function validatePassword(password: string): string {
  if (password === "") {
    return "A senha é obrigatória";
  }

  if (password.length < 6) {
    return "A senha é muito curta";
  }

  if (password.length > 255) {
    return "A senha é muito longa";
  }

  const uppercaseRegex = /[A-Z]/;
  if (!uppercaseRegex.test(password)) {
    return "A senha deve conter pelo menos uma letra maiúscula";
  }

  const lowercaseRegex = /[a-z]/;
  if (!lowercaseRegex.test(password)) {
    return "A senha deve conter pelo menos uma letra minúscula";
  }

  const numberRegex = /[0-9]/;
  if (!numberRegex.test(password)) {
    return "A senha deve conter pelo menos um número";
  }

  const specialCharacterRegex = /[^A-Za-z0-9]/;
  if (!specialCharacterRegex.test(password)) {
    return "A senha deve conter pelo menos um caractere especial";
  }

  return "";
}

export function validateName(name: string): string {
  if (name === "") {
    return "O nome é obrigatório";
  }

  if (name.length > 255) {
    return "O nome é muito longo";
  }

  if (name.length < 3) {
    return "O nome é muito curto";
  }

  const numberRegex = /[0-9]/;
  if (numberRegex.test(name)) {
    return "O nome não deve conter números";
  }

  const specialCharacterRegex = /[^A-Za-z\s]/;
  if (specialCharacterRegex.test(name)) {
    return "O nome não deve conter caracteres especiais";
  }

  return "";
}

export function validateUsername(username: string): string {
  if (username === "") {
    return "O nome de usuário é obrigatório";
  }

  if (username.length > 255) {
    return "O nome de usuário é muito longo";
  }

  if (username.length < 3) {
    return "O nome de usuário é muito curto";
  }

  const usernameRegex = /^[a-zA-Z0-9]+$/;
  if (!usernameRegex.test(username)) {
    return "O nome de usuário deve conter apenas letras e números";
  }

  return "";
}

export function validateRepeatPassword(
  password: string,
  repeatPassword: string
): string {
  if (repeatPassword === "") {
    return "A repetição da senha é obrigatória";
  }

  if (repeatPassword !== password) {
    return "As senhas não coincidem";
  }

  return "";
}