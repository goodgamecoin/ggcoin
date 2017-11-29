package mptree

import (
	"github.com/goodgamecoin/ggcoin/common"
	"github.com/goodgamecoin/ggcoin/common/rlp"
	//"github.com/rs/zerolog/log"
	"io"
)

func (n *node) EncodeRLP(w io.Writer) error {
	if err := rlp.Encode(w, n.path); err != nil {
		return err
	}
	if err := writeMap(w, n.children); err != nil {
		return err
	}
	if err := writeMap(w, n.data); err != nil {
		return err
	}
	return nil
}

func (n *node) DecodeRLP(s *rlp.Stream) error {
	path, err := s.Bytes()
	if err != nil {
		return err
	}
	children, err := readMap(s)
	if err != nil {
		return err
	}
	data, err := readMap(s)
	if err != nil {
		return err
	}
	*n = node{
		path:     path,
		children: children,
		data:     data,
	}
	return nil
}

func writeMap(w io.Writer, m map[byte]*common.Hash) error {
	data := []interface{}{}
	for i := byte(0); i < 255; i++ {
		if val := m[i]; val != nil {
			i := i
			data = append(data, &i)
			data = append(data, val)
		}
	}
	if err := rlp.Encode(w, data); err != nil {
		return err
	}
	return nil
}

func readMap(s *rlp.Stream) (map[byte]*common.Hash, error) {
	_, err := s.List()
	if err != nil {
		return nil, err
	}
	m := map[byte]*common.Hash{}
	for {
		index, err := s.Uint()
		if err != nil {
			if err == rlp.EOL {
				break
			}
			return nil, err
		}
		value, err := s.Bytes()
		if err != nil {
			return nil, err
		}
		if len(value) != common.HashLength {
			return nil, ErrBadEncodedBinary
		}
		var h common.Hash
		copy(h[:], value)
		m[byte(index)] = &h
	}
	if err := s.ListEnd(); err != nil {
		return nil, err
	}
	return m, nil
}
