using System;
using System.Runtime.InteropServices;

public static class NexusCloud
{
    // Declaração da função Go GetSQSEvents
    [DllImport("libevents.so", CallingConvention = CallingConvention.Cdecl)]
    public static extern int GetSQSEvents(out IntPtr output, out int outputLen);

    // Declaração da função Go FreeMemory
    [DllImport("libevents.so", CallingConvention = CallingConvention.Cdecl)]
    public static extern void FreeMemory(IntPtr ptr);
}
