using System;
using System.Runtime.InteropServices;
using Google.Protobuf;
using Protos; // Importa o pacote gerado do Protobuf

class Program
{
    // Importa as funções do Go
    [DllImport("cloudsdk.so", CallingConvention = CallingConvention.Cdecl)]
    public static extern int GetSQSEvents(out IntPtr output, out int outputLen);

    [DllImport("cloudsdk.so", CallingConvention = CallingConvention.Cdecl)]
    public static extern void FreeMemory(IntPtr ptr);

    static void Main()
    {
        // Variáveis para armazenar a resposta da função Go
        IntPtr outputPtr;
        int outputLen;

        // Chamar a função Go para obter os eventos SQS
        int result = GetSQSEvents(out outputPtr, out outputLen);

        if (result == 0)
        {
            // Converter o ponteiro para byte[]
            byte[] outputData = new byte[outputLen];
            Marshal.Copy(outputPtr, outputData, 0, outputLen);

            // Desserializar os eventos SQS
            MessageQueueList eventsList = MessageQueueList.Parser.ParseFrom(outputData);

            // Exibir os eventos
            foreach (var eventItem in eventsList.Messages)
            {
                 Console.WriteLine($"ID: {eventItem.Id}, Mensagem: {eventItem.MessageBody}, Provedor: {eventItem.Provider}");
            }

            // Liberar a memória alocada pelo Go
            FreeMemory(outputPtr);
        }
        else
        {
            Console.WriteLine("Erro ao chamar a função Go.");
        }
    }
}