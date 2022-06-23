package routes

import (
	"fmt"
	"strings"
	"webapp/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type usuario struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	c.Accepts("application/json")
	sess, err := store.Get(c)
	if strings.Split(c.Path(), "/")[1] == "auth" {
		return c.Next()
	}
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Não autorizado"})
	}
	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Não autorizado"})
	}
	return c.Next()
}

func Registro(c *fiber.Ctx) error {
	var data usuario
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "algo deu errado no recebimento dos dados"})
	}
	senha, bcErr := bcrypt.GenerateFromPassword([]byte(data.Senha), 14)
	if bcErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "algo deu errado do decript senha" + err.Error()})
	}
	data.Senha = string(senha)
	usuario := model.Usuario{
		Nome:  data.Nome,
		Email: data.Email,
		Senha: data.Senha,
	}
	err = model.NovoUsuario(&usuario)
	if err != nil {
		fmt.Println(usuario)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "algo deu errado na criacao do usuario " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "autorizado"})

}

func Login(c *fiber.Ctx) error {
	var data usuario
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Algo deu errado - recebendo os dados" + err.Error()})
	}
	var usuario model.Usuario
	if !model.VerificaEmail(data.Email, &usuario) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Usuário não encontrado"})
	}
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(data.Senha))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Senha incorreta"})
	}
	sess, sessErr := store.Get(c)
	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Sessão inválida - " + sessErr.Error()})
	}
	sess.Set(AUTH_KEY, true)
	sess.Set(USER_ID, usuario.ID)
	sessErr = sess.Save()
	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Erro ao salvar sessão - " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Usuario Logado"})
}

func Logout(c *fiber.Ctx) error {
	sess, sessErr := store.Get(c)
	if sessErr != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Usuário deslogado, sem sessão " + sessErr.Error()})
	}
	sessErr = sess.Destroy()
	if sessErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "algo deu errado " + sessErr.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Usuário saiu da sessão "})

}

func VerificaSessao(c *fiber.Ctx) error {
	sess, sessErr := store.Get(c)
	if sessErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "não autorizado"})
	}
	auth := sess.Get(AUTH_KEY)
	if auth != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "usuario logado (verificacao sessao) " + sessErr.Error()})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "não autorizado"})
	}
}

func ObtemUsuario(c *fiber.Ctx) error {
	sess, sessErr := store.Get(c)
	if sessErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "não autorizado (obter sessao)"})
	}
	auth := sess.Get(AUTH_KEY)
	if auth == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "não autorizado (obter auth key)"})
	}
	idUsuario := sess.Get(USER_ID)
	if idUsuario == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "não autorizado (obter id usuario)"})
	}
	var usuario model.Usuario
	usuario, sessErr = model.ObterUsuario(fmt.Sprint(idUsuario))
	usuario.Senha = ""
	if sessErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "não autorizado (obter usuario)"})
	}
	return c.Status(fiber.StatusOK).JSON(usuario)
}
