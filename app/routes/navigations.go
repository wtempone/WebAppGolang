package routes

import (
	"fmt"
	"webapp/navigators"

	"github.com/gofiber/fiber/v2"
)

func ObtemListaVoos(c *fiber.Ctx) error {
	lista, err := navigators.ObtemListaVoos()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "nao foi possivel obter lista"})
	}
	return c.Status(fiber.StatusOK).JSON(lista)
}

func ObtemDetalheVoo(c *fiber.Ctx) error {
	data := c.Params("link")
	if data == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Algo deu errado - recebendo os dados detalhes voo"})
	}
	fmt.Println(data)
	detalhes, err := navigators.ObtemDetalheVoo(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "nao foi possivel obter detalhes voo"})
	}
	return c.Status(fiber.StatusOK).JSON(detalhes)
}
