using System;
using System.Runtime.InteropServices;
using Google.Protobuf;
using Protos; // Importa o pacote gerado do Protobuf

class Program
{
    static void Main()
    {
        // Variáveis para armazenar a resposta da função Go
        IntPtr outputPtr;
        int outputLen;

        // Chamar a função Go para obter os eventos SQS
        int result = NexusCloud.GetSQSEvents(out outputPtr, out outputLen);

        if (result == 0)
        {
            // Converter o ponteiro para byte[]
            byte[] outputData = new byte[outputLen];
            Marshal.Copy(outputPtr, outputData, 0, outputLen);

            // Desserializar os eventos SQS
            SQSMessageList eventsList = SQSMessageList.Parser.ParseFrom(outputData);

            // Exibir os eventos
            foreach (var eventItem in eventsList.Messages)
            {
                Console.WriteLine($"ID: {eventItem.Id}, Message: {eventItem.MessageBody}");
            }

            // Liberar a memória alocada pelo Go
            NexusCloud.FreeMemory(outputPtr);
        }
        else
        {
            Console.WriteLine("Erro ao chamar a função Go.");
        }
    }
}
