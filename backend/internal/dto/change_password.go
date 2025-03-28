package dto

// ChangePassword representa la estructura de datos para cambiar la contraseña de un usuarios
type ChangePassword struct {
	CurrentPassword string `json:"current_password"` // Contraseña actual del usuario
	NewPassword     string `json:"new_password"`     // Contraseña nueva del usuario
	ConfirmPassword string `json:"confirm_password"` // Confirmar la contraseña nueva
}
