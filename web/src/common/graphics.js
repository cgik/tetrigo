export function drawBlock(ctx, object, game) {
    ctx.beginPath();
    ctx.rect(
        object.x * game.boardWidth / 6,
        (object.y * game.boardWidth / 6) - game.boardPosition,
        game.boardWidth / 6, game.boardHeight / 12
    );

    switch (object.color) {
        case 'blue':
            ctx.fillStyle = 'hsla(208, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(208, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(208, 100%, 40%, 1)'
            ctx.fill();
            break;
        case 'yellow':
            ctx.fillStyle = 'hsla(62, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(62, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(62, 100%, 40%, 1)'
            ctx.fill();
            break;
        case 'red':
            ctx.fillStyle = 'hsla(360, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(360, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(360, 100%, 40%, 1)'
            ctx.fill();
            break;
        case 'green':
            ctx.fillStyle = 'hsla(123, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(123, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(123, 100%, 40%, 1)'
            ctx.fill();
            break;
        case 'purple':
            ctx.fillStyle = 'hsla(274, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(274, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(274, 100%, 40%, 1)'
            ctx.fill();
            break;
        case 'teal':
            ctx.fillStyle = 'hsla(177, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(177, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(177, 100%, 40%, 1)'
            ctx.fill();
            break;
    }
}

export function drawCursor(ctx, object, game) {
    ctx.beginPath();
    ctx.strokeStyle = 'white';
    ctx.lineWidth = '3';
    ctx.rect(
        object[0] * game.boardWidth / 6,
        (object[1] * game.boardWidth / 6) - game.boardPosition,
        game.boardWidth / 3, game.boardHeight / 12);
    ctx.stroke();
    ctx.closePath();
}
