#include <SDL2/SDL.h>

#define WIDTH  640
#define HEIGHT 480

SDL_Window  *window;
SDL_Surface *primarySurface;
SDL_Surface *image;

int load() {
    SDL_Surface *tempImage;

    tempImage = SDL_LoadBMP("test1.bmp");

    if (!tempImage) {
        puts("Error loading image");
        return 0;
    }

    image = SDL_ConvertSurface(tempImage, primarySurface->format, 0);

    SDL_FreeSurface(tempImage);

    if (!image) {
        puts("Error converting surface");
        return 0;
    }
    return 1;
}

int init() {
    if (SDL_Init(SDL_INIT_VIDEO ) < 0) {
        puts("Error initializing SDL");
        return 0;
    } 

    window = SDL_CreateWindow("SDL2 example #2", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, SDL_WINDOW_SHOWN);
    if (!window) {
        puts("Error creating window");
        return 0;
    }

    primarySurface = SDL_GetWindowSurface(window);
    if (!primarySurface) {
        puts("Error getting surface");
        return 0;
    }
    return 1;
}

void destroy() {
    SDL_FreeSurface(image);

    SDL_DestroyWindow(window);
    SDL_Quit();
}

int main(int argc, char** args) {
    if (!init()) {
        return 1;
    }
    if (!load()) {
        return 1;
    }

    SDL_BlitSurface(image, NULL, primarySurface, NULL);

    SDL_UpdateWindowSurface(window);
    SDL_Delay(5000);

    destroy();
    return 0;
}

