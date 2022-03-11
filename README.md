# API-sencilla-en-GO
Repositorio de pruebas para el desarrollo de una API REST básica en GO

La idea base era enviar un objeto JSON mediante POST a la api, esta parseara el objeto, hará un request a otro endpoint (Se hizo local, pero se pensó para obtener info de otra API) y retornara el objeto ya parseado junto a la data retornada de otro endpoint con la estructura de datos definida.

Igualmente se agregó que tomara content.Cond, y revise el string, si cuenta con una vocal, este sumará los numeros A y B, si no hay vocal, este los multiplicará, retornando value dentro de cada objeto calculado. Esto fue hecho para probar la efectividad de procesamiento de datos a larga escala.  

{
    "mensaje": "Prueba con go",  
    "id": 1,  
    "key": 1234,  
    "boolean": false,  
    "string": "abcde",  
    "date": "2022-03-09T17:22:00Z",  
    "file": "aca va el base 64, pero es muy largo para las pruebas",  
    "null": null,  
    "content": [  
        {  
            "NumberA": 2,  
            "NumberB": 11,  
            "Cond": "lp9"  
        },  
        {  
            "NumberA": 3,  
            "NumberB": 14,  
            "Cond": "lap"  
        }   
    ]  
}  
