import java.nio.ByteBuffer;

public class CloudSdk {

    // Carregar a biblioteca nativa
    static {
        try {
            System.out.println("Tentando carregar a biblioteca...");
            System.loadLibrary("cloudsdk"); // Certifique-se que "cloudsdk.so" está no caminho correto
            System.out.println("Biblioteca carregada com sucesso!");
        } catch (UnsatisfiedLinkError e) {
            System.err.println("Erro ao carregar a biblioteca nativa: " + e.getMessage());
        }
    }

    // Método nativo correspondente ao Go
    public static native int GetSQSEvents(DataPointer output, IntPointer outputLen);
    public static native void FreeMemory(DataPointer ptr);

    public static void main(String[] args) {
        System.out.println("Chamando GetSQSEvents...");

        // Criando ponteiros para armazenar os valores retornados do Go
        DataPointer output = new DataPointer();
        IntPointer outputLen = new IntPointer();

        // Chamando a função nativa
        int result = GetSQSEvents(output, outputLen);

        if (result == 0) {
            System.out.println("Eventos recebidos! Tamanho: " + outputLen.getValue());

            // Converter o ponteiro para um array de bytes
            byte[] eventData = output.getData(outputLen.getValue());

            // Aqui você pode desserializar com Protobuf, por exemplo:
            // MessageQueueList messageQueueList = MessageQueueList.parseFrom(eventData);

            System.out.println("Dados recebidos: " + new String(eventData));

            // Libera a memória alocada em Go
            FreeMemory(output);
        } else {
            System.out.println("Erro ao obter eventos. Código: " + result);
        }
    }

    // Classe para gerenciar o ponteiro de saída
    public static class DataPointer {
        private long address;

        public byte[] getData(int length) {
            if (address == 0) {
                return new byte[0];
            }

            ByteBuffer buffer = ByteBuffer.allocateDirect(length);
            //buffer.put(Native.getByteArray(address, length));
            return buffer.array();
        }
    }

    // Classe para armazenar o tamanho dos dados retornados
    public static class IntPointer {
        private int value;

        public int getValue() {
            return value;
        }
    }
}
