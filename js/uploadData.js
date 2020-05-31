function primerFiltro() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = (i == 1 && j == 1) ? 1 : 0
    document.getElementById('smooth').value = 1
}

function brillo() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = (i == 1 && j == 1) ? 2 : 0
    document.getElementById('smooth').value = 1
}

function masBrillo() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = (i == 1 && j == 1) ? 3 : 0
    document.getElementById('smooth').value = 1
}

function menosBrillo() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = (i == 1 && j == 1) ? 1 : 0
    document.getElementById('smooth').value = 2
}

function pasaBajo() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = 1
    document.getElementById('smooth').value = 9
}

function pasaAlto() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = (i == 1 && j == 1) ? 1 : 8
    document.getElementById('smooth').value = 9
}

function preWitt() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = i - 1
    document.getElementById('smooth').value = 1
}

function preWittMod() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = j - 1
    document.getElementById('smooth').value = 1
}

function preWittComb() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = i + j - 2
    document.getElementById('smooth').value = 1
}

function sobel() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = (j == 1) ? i - 1 * 2 : i - 1
    document.getElementById('smooth').value = 1
}

function sobelMod() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = (i == 1) ? j - 1 * 2 : j - 1
    document.getElementById('smooth').value = 1
}

function sobelComb() {
    for (let i = 0; i < 3; i++)
        for (let j = 0; j < 3; j++)
            document.getElementById('matrix' + i + '-' + j).value = ((j == 1) ? i - 1 * 2 : i - 1) + ((i == 1) ? j -
                1 * 2 : j - 1)
    document.getElementById('smooth').value = 1
}