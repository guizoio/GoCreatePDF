$(document).ready(function () {
    pdf.eventos.init();
})
var pdf = {};

pdf.eventos = {
    init: () => {
        $("#gerar_pdf").on('click', () => {
            pdf.metodos.validar();
        });
    }
}

pdf.metodos = {
    validar: () => {
        var nome        = $("#nome").val().trim();
        var cpf         = $("#cpf").val().trim();
        var rg          = $("#rg").val().trim();
        var birth_date  = $("#nascimento").val().trim();
        var code_postal = $("#cep").val().trim();
        var address     = $("#endereco").val().trim();
        var number      = $("#numero").val().trim();
        var district    = $("#bairro").val().trim();
        var city        = $("#cidade").val().trim();
        var state       = $("#estado").val().trim();
        var email       = $("#email").val().trim();
        var cell        = $("#celular").val().trim();
        var telephone   = $("#telefone").val().trim();

        var settings = {
            "url": "http://localhost:8080/create",
            "method": "POST",
            "timeout": 0,
            "headers": {
                "Content-Type": "application/json"
            },
            "data": JSON.stringify({
                "file_pdf": 1,
                "file_img": "lim.png",
                "name": nome,
                "cpf": cpf,
                "rg": rg,
                "birth_date": birth_date,
                "code_postal": code_postal,
                "address": address,
                "number": number,
                "district": district,
                "city": city,
                "state": state,
                "email": email,
                "cell": cell,
                "telephone": telephone
            }),
        };
        $.ajax(settings).done(function (response) {
            pdf.metodos.download(response);
        });
    },

    download: (arquivo) => {
        const downloadLink = document.createElement("a");
        downloadLink.href = "http://localhost:8080/storage/download/pdf/"+arquivo
        downloadLink.download = arquivo;
        downloadLink.click();
    },
}