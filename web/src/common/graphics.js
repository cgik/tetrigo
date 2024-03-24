export function drawBlock(ctx, object, game) {
    ctx.beginPath();
    ctx.rect(
        object.x * ctx.canvas.width / game.width,
        (object.y * ctx.canvas.width / game.width) - game.position,
        ctx.canvas.width / game.width, ctx.canvas.height / game.height
    );

    switch (object.color) {
        case 0:
            ctx.fillStyle = 'hsla(208, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(208, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(208, 100%, 40%, 1)'
            ctx.fill();
            break;
        case 1:
            ctx.fillStyle = 'hsla(62, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(62, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(62, 100%, 40%, 1)'
            ctx.fill();
            break;
        case 2:
            ctx.fillStyle = 'hsla(360, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(360, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(360, 100%, 40%, 1)'
            ctx.fill();
            break;
        case 3:
            ctx.fillStyle = 'hsla(123, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(123, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(123, 100%, 40%, 1)'
            ctx.fill();
            break;
        case 4:
            ctx.fillStyle = 'hsla(274, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(274, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(274, 100%, 40%, 1)'
            ctx.fill();
            break;
        case 5:
            ctx.fillStyle = 'hsla(177, 100%, 70%, 1)'
            if (!object.active) ctx.fillStyle = 'hsla(177, 100%, 24%, 1)'
            if (object.matched) ctx.fillStyle = 'hsla(177, 100%, 40%, 1)'
            ctx.fill();
            break;
        default:
            ctx.fillStyle = 'hsla(0, 0%, 0%, 1)'
            ctx.fill();
            break;
    }

    ctx.strokeStyle = 'black';
    ctx.lineWidth = '3';
    ctx.rect(
        object.x * ctx.canvas.width / game.width,
        (object.y * ctx.canvas.width / game.width) - game.position,
        ctx.canvas.width / game.width, ctx.canvas.height / game.height
    );
    ctx.stroke();
    ctx.closePath();

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

export function renderCanvas(ctx, game) {
    ctx.clearRect(0, 0, ctx.canvas.width, ctx.canvas.height);

    game.board.blocks.forEach(
        block => drawBlock(ctx, block, game.board)
    )
}