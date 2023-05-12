package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456") // Cria o usuario
	assert.Nil(t, err)                                    // Dever ser nil pois não gerou erro.
	assert.NotNil(t, user)                                // Não pode retornar nil para o usuario
	assert.NotEmpty(t, user.ID)                           // Não pode faltar o ID
	assert.NotEmpty(t, user.Password)                     // Não pode faltar o Password
	assert.Equal(t, "John Doe", user.Name)                // Deve ser igual o nome criado
	assert.Equal(t, "j@j.com", user.Email)                // Deve ser igual o email criado
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456") // Cria o usuário
	assert.Nil(t, err)                                    // Dever ser nil pois não gerou erro.
	assert.True(t, user.ValidatePassword("123456"))       //Deve retornar verdadeiro pois o password está correto
	assert.False(t, user.ValidatePassword("1234567"))     //Deve retornar false pois o password está incorreto
	assert.NotEqual(t, "123456", user.Password)           // Não pode ser igual, pois o password deve está criptografado
}
