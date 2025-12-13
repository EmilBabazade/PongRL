package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func loadImage(path string) (*rl.Image, error) {
	image := rl.LoadImage(path)
	if image.Data == nil {
		return nil, fmt.Errorf("file not found at path %s", path)
	}
	return image, nil
}

func loadFont(path string) (rl.Font, error) {
	font := rl.LoadFont(path)
	if font.Chars == nil {
		return rl.Font{}, fmt.Errorf("file not found at path %s", path)
	}
	return font, nil
}

func loadTexture(path string) (rl.Texture2D, error) {
	texture := rl.LoadTexture(path)
	if texture.Width == 0 && texture.Height == 0 {
		return rl.Texture2D{}, fmt.Errorf("file not found at path %s", path)
	}
	return texture, nil
}

func loadSound(path string) (rl.Sound, error) {
	sound := rl.LoadSound(path)
	if sound.FrameCount == 0 {
		return rl.Sound{}, fmt.Errorf("file not found at path %s", path)
	}
	return sound, nil
}

func loadMusicStream(path string) (rl.Music, error) {
	stream := rl.LoadMusicStream(path)
	if stream.FrameCount == 0 {
		return rl.Music{}, fmt.Errorf("file not found at path %s", path)
	}
	return stream, nil
}
